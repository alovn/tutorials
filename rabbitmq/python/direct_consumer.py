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


channel.exchange_declare(exchange='direct_logs', type='direct')

# 临时队列，关闭则会删除
result = channel.queue_declare(exclusive=True)
#result = channel.queue_declare(queue='direct-logs1', durable=True)

#用于获取临时queue的name
queue_name = result.method.queue

print queue_name

# 根据 Binding 指定的 Routing Key, 将符合Key的消息发送到 Binding 的 Queue (接收者只处理binding和routingKey相同的消息)
#若多个消费者共用一个queue,则会轮询消费
#若多个消费者用不同的queue,则会广播发送
channel.queue_bind(exchange='direct_logs', queue=queue_name, routing_key="log1")


def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)


channel.basic_consume(callback,
                      queue=queue_name,
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()