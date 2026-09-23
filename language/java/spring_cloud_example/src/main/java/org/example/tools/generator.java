package org.example.tools;

import com.baomidou.mybatisplus.annotation.FieldFill;
import com.baomidou.mybatisplus.generator.FastAutoGenerator;
import com.baomidou.mybatisplus.generator.config.OutputFile;
import com.baomidou.mybatisplus.generator.engine.FreemarkerTemplateEngine;
import com.baomidou.mybatisplus.generator.fill.Column;

import java.nio.file.Paths;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class generator {
    public static void main(String[] args) {
        String basePath = Paths.get("src/main").toAbsolutePath().toString();
        System.out.println(basePath);
        FastAutoGenerator.create("jdbc:postgresql://localhost:5432/janus", "janus", "janus")
                // 全局配置
                .globalConfig((scanner, builder) -> builder
                        .author("janus")
                        .outputDir(basePath + "/java"))
                // 包配置
                .packageConfig((scanner, builder) -> builder
                        .parent("org.example")
                        .pathInfo(Collections.singletonMap(OutputFile.xml, basePath + "/resources/mapper")))
                // 策略配置
                .strategyConfig((scanner, builder) -> builder.addInclude(getTables(scanner.apply("请输入表名，多个英文逗号分隔？所有输入 all"))).entityBuilder().enableLombok().addTableFills(new Column("create_time", FieldFill.INSERT)).build())
                // 使用Freemarker引擎模板，默认的是Velocity引擎模板
                .templateEngine(new FreemarkerTemplateEngine())
                .execute();
    }

    // 处理 all 情况
    protected static List<String> getTables(String tables) {
        return "all".equals(tables) ? Collections.emptyList() : Arrays.asList(tables.split(","));
    }
}
