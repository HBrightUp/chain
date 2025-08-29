#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "coinview.h"
#include <QFileDialog>
#include "json.hpp"
#include "walletdb.h"
#include <QMessageBox>
using json = nlohmann::json;
#include <sstream>
#include <fstream>
#include "rpc.h"
#include "dialogtokenbranch.h"

WalletDB MainWindow::wallet_;
static std::map<size_t, WalletDB::EthToken> s_map_index_ethtoken;
static std::map<size_t, WalletDB::Branch> s_map_index_branch;

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    open_wallet_ = false;
    //    refreshFee();
}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::on_btn_eth_submit_clicked()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"wallet","please open config file");
        return;
    }

    QString from = ui->lineEdit_eth_from->text();
    QString to = ui->lineEdit_eth_to->text();
    QString amount = ui->lineEdit_eth_amount->text();
    QString fee = ui->lineEdit_eth_fee->text();
    QString privkey ;
    bool ret = false;

    if(map_address_privkey_.find(from) != map_address_privkey_.end())
    {
        privkey = map_address_privkey_[from];
    }
    else
    {
        ret = wallet_.getPrivkey(from, privkey);
        if(!ret)
        {
            return;
        }
        map_address_privkey_[from] = privkey;
    }

    HDwallet::Transfer transfer;
    size_t index = ui->comboBox_token->currentIndex();
    transfer.coin_type = HDwallet::ETH;
    transfer.amount = amount.toStdString();
    transfer.fee = fee.toStdString();
    transfer.to = to.toStdString();
    transfer.from = from.toStdString();

    if (index > 0)
    {
        WalletDB::EthToken token = s_map_index_ethtoken[index];
        transfer.is_token = true;
        transfer.contract = token.contract_address.toStdString();
        transfer.decimal = token.decimal;
    }

    std::string txid;
    ret = hd_wallet_.transferCoin(transfer,txid);

    if(ret)
    {
         QMessageBox::warning(this,"txid",txid.c_str());
         qInfo(txid.c_str());
         WalletDB::Tx tx;
         tx.amount = transfer.amount.c_str();
         tx.height = 0;
         tx.to = transfer.to.c_str();
         tx.txid = txid.c_str();
         tx.from = transfer.from.c_str();
         tx.coin = "ETH";
         if (transfer.is_token)
         {
             tx.coin = ui->comboBox_token->currentText();
         }
         wallet_.addTx(tx);
    }
    else
    {
        QMessageBox::warning(this, "error", "transaction fail");
    }

    ui->lineEdit_eth_from->clear();
    ui->lineEdit_eth_to->clear();
    ui->lineEdit_eth_amount->clear();
    ui->lineEdit_eth_fee->clear();
    refreshFee();
}

void MainWindow::on_btn_btc_submit_clicked()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"wallet","please open config file");
        return;
    }

    HDwallet::Transfer transfer;
    QString from = ui->lineEdit_btc_from->text();
    transfer.coin_type = HDwallet::BTC;
    transfer.from = from.toStdString();
    transfer.to = ui->lineEdit_btc_to->text().toStdString();
    transfer.amount = ui->lineEdit_btc_amount->text().toStdString();
    transfer.fee = ui->lineEdit_btc_fee->text().toStdString();
    QString privkey ;
    bool ret = false;

    if(map_address_privkey_.find(from) != map_address_privkey_.end())
    {
        privkey = map_address_privkey_[from];
    }
    else
    {
        ret = wallet_.getPrivkey(from, privkey);
        if(!ret)
        {
            return;
        }
        map_address_privkey_[from] = privkey;
    }
    size_t index = ui->comboBox_branch->currentIndex();
    if (index > 0)
    {
        WalletDB::Branch branch = s_map_index_branch[index];
    }

    std::string txid;
    ret = hd_wallet_.transferCoin(transfer,txid);

    if(ret)
    {
         QMessageBox::warning(this,"txid",txid.c_str());
         qInfo(txid.c_str());
         WalletDB::Tx tx;
         tx.amount = transfer.amount.c_str();
         tx.height = 0;
         tx.to = transfer.to.c_str();
         tx.txid = txid.c_str();
         tx.from = transfer.from.c_str();
         tx.coin = "BTC";
         if (index > 0)
         {
             tx.coin = ui->comboBox_branch->currentText();
         }
         wallet_.addTx(tx);
    }
    else
    {
        QMessageBox::warning(this, "error", "transaction fail");
    }

    ui->lineEdit_btc_from->clear();
    ui->lineEdit_btc_to->clear();
    ui->lineEdit_btc_amount->clear();
    ui->lineEdit_btc_fee->clear();
    refreshFee();
}

void MainWindow::on_actiontx_triggered()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"Title","please open config file");
        return;
    }
    wallet_.open();
    view_table_.show();
    wallet_.close();
    refreshFee();
}

void MainWindow::on_actiongenerate_triggered()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"Title","please open config file");
        return;
    }

    int current = ui->tabcoin->currentIndex();
    wallet_.open();
    std::string url;
    if(current == 0)
        Rpc::instance().getEthUrl(url);
    else if(current == 1)
        Rpc::instance().getBtcUrl(url);

    conf_dlg_.setUrl(QString(url.c_str()));
    conf_dlg_.show();
    wallet_.close();
    refreshFee();
}

void MainWindow::on_actionopen_triggered()
{
    QString fileName = QFileDialog::getOpenFileName(
                this,
                tr("open a file."),
                "./",
                tr("config(*.json);;All files(*.*)"));

    if(fileName.isEmpty())
        return;
    std::ifstream jfile(fileName.toStdString().c_str());
    if(!jfile)
    {
        jfile.close();
    }
    json json_conf;
    jfile >> json_conf;
    if(!json_conf.is_object())
    {
        jfile.close();
    }
    jfile.close();

    std::string db_path = json_conf["path"].get<std::string>();
    std::string db_name = json_conf["db"].get<std::string>();
    std::string eth_url = json_conf["eth"].get<std::string>();
    std::string btc_url = json_conf["btc"]["url"].get<std::string>();
    std::string btc_auth = json_conf["btc"]["auth"].get<std::string>();
    Rpc::instance().initInstance(btc_url, btc_auth, eth_url);

    wallet_.init(db_path,db_name);
    config_file_ = fileName;
    open_wallet_ = true;
    std::vector<WalletDB::EthToken> vect_token;
    wallet_.getEthTokenInfo(vect_token);
    ui->comboBox_token->insertItem(0, "");
    for (size_t i = 0; i < vect_token.size(); i++)
    {
        s_map_index_ethtoken[i+1] = vect_token[i];
        ui->comboBox_token->insertItem(i+1, vect_token[i].name);
    }

    std::vector<WalletDB::Branch> vect_branch;
    wallet_.getBtcBranch(vect_branch);
    ui->comboBox_branch->insertItem(0, "BTC");
    for (size_t i = 0; i < vect_branch.size(); i++)
    {
         s_map_index_branch[i+1] = vect_branch[i];
         ui->comboBox_branch->insertItem(i+1, vect_branch[i].name);
    }
   // wallet_.close();

    refreshFee();
}

void MainWindow::on_actionbalance_triggered()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"Title","please open config file");
        return;
    }
    wallet_.open();
    view_balance_.resetModel();
    view_balance_.show();
    wallet_.close();
    refreshFee();

}

void MainWindow::refreshFee()
{
    setFeeEth();
    setFeeBtc();
}

void MainWindow::setFeeEth()
{
    size_t index = ui->comboBox_token->currentIndex();
    uint64_t gas_limit ;
    std::string gas_price ;
    Rpc::instance().getGasPrice(gas_price);

    if(index > 0)
    {
        gas_limit = 60000;
    }
    else
    {
        gas_limit = 21000;
    }

    //std::cout << "gas price: " << gas_price << std::endl;
    uint64_t fee = gas_limit * QString(gas_price.c_str()).toLong(nullptr,16);
    double fee_eth = (double)fee / 1000000000.0 / 1000000000.0;
    //std::cout << fee_eth << std::endl;
    ui->lineEdit_eth_fee->setText(std::to_string(fee_eth).c_str());

}

void MainWindow::setFeeBtc()
{
    QString fee = "0.00001";
    ui->lineEdit_btc_fee->setText(fee);
}

void MainWindow::on_btn_eth_open_clicked()
{
    QString address;
//    setAddressPrivkey(address);
    ui->lineEdit_eth_from->setText(address);
}

void MainWindow::setAddressPrivkey(/*QString& address*/)
{
    QString fileName = QFileDialog::getOpenFileName(
                this,
                tr("open a wallet file."),
                "D:/",
                tr("config(*.json);;All files(*.*)"));

    if (fileName.isEmpty())
    {
        return;
    }
    pass_dlg_.setConfirm(false);
    pass_dlg_.exec();

    //qInfo(fileName.toStdString().c_str());
    QString password = pass_dlg_.getPassword();
 /*   Account account;
    if (!ETHAPI::getInstance().decryptKeystore(fileName.toStdString(), password.toStdString(), account))
    {
        return;
    }
    qWarning(account.privateKey.c_str());
    qWarning(account.address.c_str());
    address = account.address.c_str();
    map_address_privkey_[address] = QString(account.privateKey.c_str());*/
}



void MainWindow::on_btn_btc_insert_clicked()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"wallet","please open config file");
        return;
    }

    DialogTokenBranch dlg;
    dlg.activeBranch();
    WalletDB::Branch branch;
    dlg.exec();
    dlg.getBranch(branch);
    if(branch.auth.isEmpty() ||
       branch.name.isEmpty() ||
       branch.url.isEmpty()  )
    {
        qInfo("branch is empty");
        return ;
    }

    size_t index = s_map_index_branch.size() + 1;
    s_map_index_branch[index] = branch;
    ui->comboBox_token->insertItem(index, branch.name);
    wallet_.addBranch(branch);

}

void MainWindow::on_btn_eth_insert_clicked()
{
    if (!open_wallet_)
    {
        QMessageBox::warning(this,"wallet","please open config file");
        return;
    }

    DialogTokenBranch dlg;
    WalletDB::EthToken token;
    dlg.exec();
    dlg.getToken(token);
    if(token.contract_address.isEmpty() ||
       token.name.isEmpty() ||
       token.decimal <= 0   )
    {
        qInfo("token is empty");
        return ;
    }

    size_t index = s_map_index_ethtoken.size() + 1;
    s_map_index_ethtoken[index] = token;
    ui->comboBox_token->insertItem(index, token.name);
    wallet_.addEthToken(token);
}

void MainWindow::on_comboBox_token_currentIndexChanged(int index)
{
    setFeeEth();
}
