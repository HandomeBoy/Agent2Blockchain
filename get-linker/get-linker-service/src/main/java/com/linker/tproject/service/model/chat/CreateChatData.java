package com.linker.tproject.service.model.chat;

import lombok.Data;

@Data
public class CreateChatData {
    private String id;
    private String bot_id;
    private Integer completed_at;
    private String conversation_id;

}
