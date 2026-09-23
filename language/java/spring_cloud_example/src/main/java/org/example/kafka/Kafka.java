package org.example.kafka;

import com.fasterxml.jackson.databind.deser.std.StringDeserializer;
import org.apache.kafka.clients.admin.AdminClient;
import org.apache.kafka.clients.admin.CreateTopicsResult;
import org.apache.kafka.clients.admin.NewTopic;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.KafkaFuture;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Configurable;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.DefaultKafkaConsumerFactory;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import java.time.Duration;
import java.util.Collections;
import java.util.Properties;
import java.util.concurrent.ExecutionException;

//@Component
public class Kafka {
    @Value("${spring.kafka.bootstrap-servers}")
    private String bootstrapServers;

    private static final String topic = "example-topic";

    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    public void creatTopic() {
        Properties config = new Properties();
        config.put("bootstrap.servers", bootstrapServers);

        try (AdminClient adminClient = AdminClient.create(config)) {
            NewTopic newTopic = new NewTopic(topic, 1, (short) 1); // topic名称，分区数，副本因子

            CreateTopicsResult result = adminClient.createTopics(Collections.singletonList(newTopic));

            KafkaFuture<Void> future = result.all();
            future.get(); // 这里会阻塞直到操作完成

            System.out.println("Topic created successfully.");
        } catch (InterruptedException | ExecutionException e) {
            System.out.println("Failed to create topic.");
            e.printStackTrace();
        }
    }

    public void sendMessage(String message) {
        kafkaTemplate.send(topic, message);
    }

}
