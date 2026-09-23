package org.example.tasks;

import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

@Service
public class ScheduledTasks {

    // 每隔5秒执行一次
    @Scheduled(fixedRate = 5000)
    public void reportCurrentTimeWithFixedRate() {
        System.out.println("Fixed Rate Task :: Execution Time - " + System.currentTimeMillis());
        throw new RuntimeException("This is a scheduled task test");
    }
}