package com.linker.tproject.service.config;

import lombok.Getter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Getter
@Configuration
public class CozeConfig {

    // getter 方法
    @Value("${coze.http.createConversation}")
    private String cozeConservation;

    @Value("${coze.create.agent_id}")
    private String cozeAgentId;

    @Value("${coze.bearer.personal}")
    private String cozeBearer;

    @Value("${coze.http.createMessage}")
    private String cozeMessage;

    @Value("${coze.http.uploadFile}")
    private String cozeUploadFile;

    @Value("${tomato.agent.img.path}")
    private String imgRootPath;

    @Value("${coze.http.getFile}")
    private String cozeGetFile;

    @Value("${coze.http.createChat}")
    private String cozeChat;

    @Value("${coze.http.getChatMessageList}")
    private String cozeGetAgentMessageList;

    @Value("${coze.http.getRetrieveChat}")
    private String cozeRetrieveChat;

    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }

}