# encoding:utf-8
import pika

credentials = pika.PlainCredentials('guest', 'guest')
connection = pika.BlockingConnection(pika.ConnectionParameters(
	host='s1004.lab.org',
	port=5672,
	virtual_host='/',
	credentials=credentials))
channel = connection.channel()

# durable=True声明queue是持久化的，这样即便Rabb崩溃了重启后queue仍然存在
channel.queue_declare(queue='hello', durable=True)

# 除了要声明queue是持久化的外，还需声明message是持久化的
# basic_publish的properties参数指定message的属性
# 此处pika.BasicProperties中的delivery_mode=2指明message为持久的
# 这样一来RabbitMQ崩溃重启后queue仍然存在其中的message也仍然存在
# 需注意的是将message标记为持久的并不能完全保证message不丢失，因为
# 从RabbitMQ接收到message到将其存储到disk仍需一段时间，若此时RabbitMQ崩溃则message会丢失
# 况且RabbitMQ不会对每条message做fsync动作
# 可通过publisher confirms实现更强壮的持久性保证
channel.basic_publish(exchange='',
                      routing_key='hello',
                      properties=pika.BasicProperties(
                          delivery_mode=2,
                      ),
                      body='Hello World!')
print(" [x] Sent 'Hello World!'")
connection.close()
