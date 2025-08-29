#include "dialogmasterkey.h"
#include "ui_dialogmasterkey.h"

DialogMasterKey::DialogMasterKey(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::DialogMasterKey)
{
    ui->setupUi(this);
}

DialogMasterKey::~DialogMasterKey()
{
    delete ui;
}

void DialogMasterKey::setShowText(const QString &privkey, const QString &pubkey, const QString &mnemonic)
{
    ui->textEdit_privkey->setText(privkey);
    ui->textEdit_pubkey->setText(pubkey);
    ui->textEdit->setText(mnemonic);
}
