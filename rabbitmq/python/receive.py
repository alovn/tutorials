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


channel.queue_declare(queue='hello', durable=True)


def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
    print method.delivery_tag
    time.sleep(2)
    print 'Done'
    # 对message进行确认
    #ch.basic_ack(delivery_tag=method.delivery_tag)

#若存在多个consumer每个consumer的负载可能不同，有些处理的快有些处理的慢
#RabbitMQ并不管这些，只是简单的以round-robin的方式分配message
#这可能造成某些consumer积压很多任务处理不完而一些consumer长期处于饥饿状态
#可以使用prefetch_count=1的basic_qos方法可告知RabbitMQ只有在consumer处理并确认了上一个message后才分配新的message给他
#否则分给另一个空闲的consumer
#通过配置 prefetch_count 参数, 来设置一次从队列中取多少条消息
channel.basic_qos(prefetch_count=1)

#没有设置no_ack=True则consumer在收到message后会向RabbitMQ反馈已收到并处理了message告诉RabbitMQ可以删除该message
channel.basic_consume(callback,
                      queue='hello',
                      no_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()