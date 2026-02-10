package com.linker.tproject.service.req.message;

import lombok.Data;

@Data
public class CreateMessageRequest {
    private String role;
    private String content; // 改为 List<MessageContent>
    private String content_type;
}
