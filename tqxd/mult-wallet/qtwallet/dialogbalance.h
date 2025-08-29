#ifndef DIALOGBALANCE_H
#define DIALOGBALANCE_H

#include <QDialog>
#include <QtSql>
namespace Ui {
class DialogBalance;
}

class DialogBalance : public QDialog
{
    Q_OBJECT

public:
    explicit DialogBalance(QWidget *parent = nullptr);
    ~DialogBalance();

    void resetModel();

private slots:
    void on_buttonBox_accepted();
protected:
    void updateBalance();

    void updateBtcBalance();

    void updateEthBalance();
private:
    Ui::DialogBalance *ui;
    QSqlTableModel* model_;
};

#endif // DIALOGBALANCE_H
