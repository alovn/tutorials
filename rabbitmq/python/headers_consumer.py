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

# 临时队列，关闭则会删除
result = channel.queue_declare(exclusive=True)

#用于获取临时queue的name
queue_name = result.method.queue

print queue_name

channel.queue_bind(exchange='headers_logs',
				   queue=queue_name,
				   routing_key="headers_log1",
				   #arguments={'a': '2', 'x-match': 'all'}
				   #arguments={'a': '2', 'b':'3', 'x-match': 'all'},
				   arguments={'a': '2', 'b':'3', 'x-match': 'any'}
				   )


def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)


channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()