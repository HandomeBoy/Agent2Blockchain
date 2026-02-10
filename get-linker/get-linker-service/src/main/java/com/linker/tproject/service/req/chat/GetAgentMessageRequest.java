package com.linker.tproject.service.req.chat;

import lombok.Data;

@Data
public class GetAgentMessageRequest {
    private String conversation_id;
    private String chat_id;
}
