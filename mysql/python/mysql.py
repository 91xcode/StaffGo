# -*- encoding=utf-8 -*-

# 执行 python2.7 mysql.py
import mysqldao
import MySQLdb

class Staff(object):


    def __init__(self):
        self.db_extra = self.test_get_dao()


    def test_get_dao(self,env='test'):
        """
         连接数据库
        :return:
        """
        if env=='test':
            manual_config = {'db_host': '127.0.0.1', 'db_user': 'root', 'db_pwd': '', 'db_name': 'demo',
                             'db_charset': 'utf8'}



        manual_dao = mysqldao.MysqlDao(manual_config)
        return manual_dao

    def getAll(self):
        sql = "select * from user"
        return self.db_extra.db_query(sql)

if __name__ == '__main__':
    model = Staff()
    result = model.getAll()
    print (result)