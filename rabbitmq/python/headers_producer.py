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

channel.exchange_declare(exchange='headers_logs', type='headers')


message = 'Hello, World!'
channel.basic_publish(exchange='headers_logs',
					  routing_key='headers_log1',
					  properties=pika.BasicProperties(
						  delivery_mode=2, #消息持久化
						  headers={'a': '2'}
					  ),
					  body=message)

print " [x] Sent %r" % (message,)

connection.close()