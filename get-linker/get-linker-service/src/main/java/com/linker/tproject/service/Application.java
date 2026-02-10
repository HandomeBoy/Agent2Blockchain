package com.linker.tproject.service;

import com.linker.auth.api.token.rpc.ITokenAuthRpcController;
import com.linker.tproject.api.rpc.ITestRpcController;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;
import org.springframework.context.annotation.Import;

/**
 * @desc Boot应用的入口类
 * @author lyj
 */
@SpringBootApplication(scanBasePackages = {"com.linker","com.linker.core"})
@EnableFeignClients(clients = {ITokenAuthRpcController.class})  //引入第三方rpc接口，
@MapperScan(basePackages = "com.linker.tproject.service.mapper")
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}