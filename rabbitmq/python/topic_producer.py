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

channel.exchange_declare(exchange='topic_logs', type='topic')


message = 'Hello, World!'
#将message publish到名为log的exchange中
#因为是fanout类型的exchange，这里无需指定routing_key
channel.basic_publish(exchange='topic_logs',
                      routing_key='topic.logs.info',
                      body=message)

print " [x] Sent %r" % (message,)

connection.close()