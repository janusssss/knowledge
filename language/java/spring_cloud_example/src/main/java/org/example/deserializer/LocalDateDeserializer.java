package org.example.deserializer;

import com.alibaba.fastjson.parser.DefaultJSONParser;
import com.alibaba.fastjson.parser.JSONLexer;
import com.alibaba.fastjson.parser.JSONToken;
import com.alibaba.fastjson.parser.deserializer.ObjectDeserializer;
import com.alibaba.fastjson.JSONException;

import java.lang.reflect.Type;
import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;
import java.time.OffsetDateTime;

public class LocalDateDeserializer implements ObjectDeserializer {
    // 支持毫秒部分和时区信息的格式
    private static final DateTimeFormatter DATE_TIME_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ss.SSSX");
    // 简单日期格式
    private static final DateTimeFormatter SIMPLE_DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-M-d");
    // 标准日期格式
    private static final DateTimeFormatter STANDARD_DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd");


    @Override
    public LocalDate deserialze(DefaultJSONParser parser, Type type, Object fieldName) {
        JSONLexer lexer = parser.getLexer();
        if (lexer.token() == JSONToken.LITERAL_STRING) {
            String dateStr = lexer.stringVal().replace("\"", ""); // 去除可能的引号

            // 检查是否为 null 或空字符串
            if (dateStr.trim().isEmpty()) {
                return null; // 或者根据业务需求返回一个默认值
            }

            try {
                // 尝试以标准日期格式解析
                return LocalDate.parse(dateStr, STANDARD_DATE_FORMATTER);
            } catch (DateTimeParseException e1) {
                try {
                    // 如果标准日期格式失败，尝试简单日期格式解析
                    return LocalDate.parse(dateStr, SIMPLE_DATE_FORMATTER);
                } catch (DateTimeParseException e2) {
                    try {
                        // 如果简单日期格式也失败，尝试带有时间的日期格式解析
                        return OffsetDateTime.parse(dateStr, DATE_TIME_FORMATTER).toLocalDate();
                    } catch (DateTimeParseException e3) {
                        throw new JSONException("Unable to parse date: " + dateStr, e3);
                    }
                }
            }
        } else {
            throw new JSONException("Unsupported token type for LocalDate parsing");
        }
    }

    @Override
    public int getFastMatchToken() {
        return JSONToken.LITERAL_STRING;
    }
}