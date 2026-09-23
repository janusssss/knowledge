package org.example.kafka;

import org.junit.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.AbstractJUnit4SpringContextTests;

@SpringBootTest
public class KafkaTest extends AbstractJUnit4SpringContextTests {
    @Autowired
    private Kafka kafka;

    @Test
    public void testCreatTopic() {
        kafka.creatTopic();
    }

    @Test
    public void testSendMessage() {
        String message = "Hello World";
        kafka.sendMessage(message);
    }

    @Test
    public void testConsumeMessage() {
//        kafka.consumeMessages();
    }
}
