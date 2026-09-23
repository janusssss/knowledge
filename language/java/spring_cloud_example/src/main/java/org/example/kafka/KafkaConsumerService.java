package org.example.kafka;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

//@Component
public class KafkaConsumerService {
    @KafkaListener(topics = "example-topic", groupId = "myGroup")
    public void listen(ConsumerRecord<?, ?> record) {
        System.out.println(record.value());
    }
}