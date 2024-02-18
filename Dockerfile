FROM rabbitmq:3.12.13-management-alpine
RUN rabbitmq-plugins enable --offline rabbitmq_mqtt rabbitmq_web_mqtt