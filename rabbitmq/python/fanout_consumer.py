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


#作为好的习惯，在producer和consumer中分别声明一次以保证所要使用的exchange存在
channel.exchange_declare(exchange='logs', type='fanout')

#在不同的producer和consumer间共享queue时指明queue的name是重要的
#但某些时候，比如日志系统，需要接收所有的log message而非一个子集
#而且仅对当前的message 流感兴趣，对于过时的message不感兴趣，那么
#可以申请一个临时队列这样，每次连接到RabbitMQ时会以一个随机的名字生成
#一个新的空的queue，将exclusive置为True，这样在consumer从RabbitMQ断开后会删除该queue
result = channel.queue_declare(exclusive=True)

#若为持久化的queue, 则每次发送消息都会转发到该队列，消费者停止后仍然会发送消息到queue
#若多个消费者共用一个queue,则会轮询消费
#若多个消费者用不同的queue,则会广播发送
#result = channel.queue_declare(queue='hello-logs1', durable=True)

#用于获取临时queue的name
queue_name = result.method.queue

print queue_name

channel.queue_bind(exchange='logs', queue=queue_name)


def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)


channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()