package org.example;

import org.example.config.HiveConfig;
import org.junit.Test;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class hiveTest {

    @Test
    public void selectTest() {
        Statement statement = HiveConfig.getStatement();

        String sql = "SELECT * FROM example.animal";
        try {
            ResultSet resultSet = statement.executeQuery(sql);
            System.out.println(resultSet);
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
    }
}
