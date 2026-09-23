package org.example.component.stream;

import org.junit.Test;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

@SpringBootTest
public class StreamTest {
    List<Integer> list = Arrays.asList(1, 2, 3, 3);

    @Test
    public void testFilter() {
        long count = list.stream().filter(v -> v % 2 == 0).count();
        System.out.println(count);
    }

    @Test
    public void testMap() {
        System.out.println("old: " + list);
        List<Integer> listTemp = list.stream().map(v -> v + 10).collect(Collectors.toList());
        System.out.println("new: " + listTemp);
        System.out.println("old: " + list);
    }

    @Test
    public void testPeek() {
        System.out.println("old: " + list);
        List<Integer> listTemp = list.stream().peek(v -> System.out.println(v + 10)).collect(Collectors.toList());
        System.out.println("new: " + listTemp);
    }

    @Test
    public void testCollect() {
        System.out.println("old: " + list);
        Set<Integer> set = list.stream().map(v -> v + 10).collect(Collectors.toSet());
        System.out.println("set: " + set);
        Map<Integer, Integer> map = list.stream().distinct().collect(Collectors.toMap(v -> v, v -> v + 10));
        System.out.println("map: " + map);
    }
}
