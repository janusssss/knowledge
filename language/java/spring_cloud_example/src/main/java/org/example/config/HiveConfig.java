package org.example.config;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

public class HiveConfig {
    private static Connection conn;
    private static Statement stmt;


    public static Statement getStatement() {
        String url = "jdbc:hive2://localhost:10000/example";
        String user = "janus";
        String password = "";
        try {
            conn = DriverManager.getConnection(url, user, password);
            stmt = conn.createStatement();
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
        return stmt;
    }

    public static void close() {
        try {
            conn.close();
            stmt.close();
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
    }
}
