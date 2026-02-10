package com.linker.tproject.service.req.conversation;

import lombok.Data;

@Data
public class CreateConversationRequest {
    private String bot_id;
    private String name;

    public CreateConversationRequest(String bot_id, String name) {
        this.bot_id = bot_id;
        this.name = name;
    }
}


