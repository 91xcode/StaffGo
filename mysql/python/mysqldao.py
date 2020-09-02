# coding: utf-8
"""

"""
import MySQLdb
import traceback
import logging


class MysqlDao:
    def __init__(self,config):
        self.config = config
        self.conn = None
        #self.conn = MySQLdb.connect(host=config['db_host'], user=config['db_user'], passwd=config['db_pwd'], db=config['db_name'],charset=config.get('db_charset','utf8mb4'),unix_socket='/data/mysql/mysql.sock')
        self.conn = MySQLdb.connect(host=config['db_host'], user=config['db_user'], passwd=config['db_pwd'], db=config['db_name'],charset=config.get('db_charset','utf8mb4'))
        self.conn.autocommit(True)

    def __del__(self):
        self.db_close()

    def db_close(self):
        try:
            if self.conn:
                self.conn.close()
        except:
            pass

    def db_connect(self):
        self.db_close()
        self.conn = MySQLdb.connect(host=self.config['db_host'], user=self.config['db_user'], passwd=self.config['db_pwd'], db=self.config['db_name'],charset=self.config.get('db_charset','utf8mb4'))

    def _db_execute(self,sql):
        try:
            cur = self.conn.cursor()
            cur.execute(sql)
            cur.close()
            self.conn.commit()
        except (AttributeError, MySQLdb.OperationalError):
            self.db_connect()
            cur = self.conn.cursor()
            cur.execute(sql)
            cur.close()
            self.conn.commit()


    def _db_execute_many(self,sql,sql_args):
        try:
            cur = self.conn.cursor()
            cur.executemany(sql,sql_args)
            cur.close()
            self.conn.commit()
        except (AttributeError, MySQLdb.OperationalError):
            self.db_connect()
            cur = self.conn.cursor()
            cur.execute(sql)
            cur.close()
            self.conn.commit()


    def db_execute(self,sql):
        try:
            return self._db_execute(sql)
        except MySQLdb.Error, e:
            self.conn.rollback()
            try:
                logging.getLogger().error("MySQL Error [%d]: %s, sql=%s" % (e.args[0], e.args[1], sql))
            except IndexError:
                logging.getLogger().error("MySQL Error: %s, sql=%s" % (str(e), sql))

        except:
            logging.getLogger().error("MysqlDao exception:%s" % traceback.format_exc())

        return None

    def _db_query(self,sql):
        try:
            cur = self.conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)
            cur.execute(sql)
            results = cur.fetchall()
            cur.close()
            return results
        except (AttributeError, MySQLdb.OperationalError):
            self.db_connect()
            cur = self.conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)
            cur.execute(sql)
            results = cur.fetchall()
            cur.close()
            return results

    def db_query(self,sql):
        try:
            return self._db_query(sql)
        except MySQLdb.Error, e:
            try:
                logging.getLogger().error("MySQL Error [%d]: %s" % (e.args[0], e.args[1]))
            except IndexError:
                logging.getLogger().error("MySQL Error: %s" % str(e))

        except:
            logging.getLogger().error("MysqlDao exception:%s" % traceback.format_exc())
        return None