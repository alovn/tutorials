# encoding:utf-8
import pika
import time

credentials = pika.PlainCredentials('guest', 'guest')
connection = pika.BlockingConnection(pika.ConnectionParameters(
    host='s1004.lab.org',
    port=5672,
    virtual_host='/',
    credentials=credentials))
channel = connection.channel()

#producer只能通过exchange将message发给queue
#exchange的类型决定将message路由至哪些queue
#可用的exchange类型：direct\topic\headers\fanout
#此处定义一个名称为'logs'的'fanout'类型的exchange，'fanout'类型的exchange简单的将message广播到它所知道的所有queue
channel.exchange_declare(exchange='logs', type='fanout')


message = 'Hello, World!'
#将message publish到名为log的exchange中
#因为是fanout类型的exchange，这里无需指定routing_key
channel.basic_publish(exchange='logs',
                      routing_key='',
                      body=message)

print " [x] Sent %r" % (message,)

connection.close()