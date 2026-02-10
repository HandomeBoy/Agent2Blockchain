package com.linker.tproject.service.model.message;

import lombok.Data;

@Data
public class MessageData {
    private String bot_id;
    private String chat_id;
    private String content;
    private String content_type;
    private String conversation_id;
    private Integer created_at;
    private String id;
    private Object meta_data;
    private String reasoning_content;
    private String role;
    private String section_id;
    private String type;
    private Integer updated_at;
}
