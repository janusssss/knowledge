package org.example.service;

import com.baomidou.mybatisplus.core.metadata.IPage;
import com.baomidou.mybatisplus.core.metadata.OrderItem;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.baomidou.mybatisplus.extension.toolkit.ChainWrappers;
import org.example.entity.Person;
import org.example.mapper.PersonMapper;
import org.junit.Test;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.AbstractJUnit4SpringContextTests;

import java.util.ArrayList;
import java.util.List;

@SpringBootTest
public class IPersonServiceTest extends AbstractJUnit4SpringContextTests {
    @Autowired
    private IPersonService personService;
    @Autowired
    private PersonMapper personMapper;

    @Test
    public void testSave() {
        Person person = new Person(5L, "Billie", 24, "test5@baomidou.com");
        personService.save(person);
    }

    @Test
    public void testRemove() {
        ChainWrappers.lambdaUpdateChain(personMapper)
                .eq(Person::getAge, 24)
                .remove();
    }

    @Test
    public void testList() {
        List<Person> list = personService.list();
        System.out.println(list);
    }

    @Test
    public void testLambdaChain() {
        ChainWrappers.lambdaQueryChain(Person.class)
                .eq(Person::getAge, 24)
                .list()
                .forEach(System.out::println);
    }

    @Test
    public void testSaveBatch() {
        List<Person> people = new ArrayList<>();
        people.add(new Person(100L, "Billie", 24, "test5@baomidou.com"));
        for (int i = 1; i < 100000; i++) {
            Person person = new Person();
            BeanUtils.copyProperties(people.get(0), person);
            person.setId(person.getId() + i);
            people.add(person);
        }
        personService.saveBatch(people);
    }

    @Test
    public void testPage() {
        Page<Person> page = new Page<>(1, 2);
        ChainWrappers.lambdaQueryChain(Person.class)
                .orderByAsc(Person::getAge)
                .page(page);
        System.out.println(page);
    }
}
