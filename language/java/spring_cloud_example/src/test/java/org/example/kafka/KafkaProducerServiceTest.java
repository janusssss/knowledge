package org.example.kafka;

import org.junit.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.AbstractJUnit4SpringContextTests;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.List;

@SpringBootTest
public class KafkaProducerServiceTest extends AbstractJUnit4SpringContextTests {
    @Autowired
    private KafkaProducerService producerService;

    @Test
    public void producerTest() {
        producerService.sendMessage("Hello, Kafka!");
    }

    @Test
    public void test() {
        System.out.println("Hello Kafka!");
    }
}
