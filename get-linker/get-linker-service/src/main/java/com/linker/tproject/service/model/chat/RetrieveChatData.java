package com.linker.tproject.service.model.chat;

import cn.hutool.json.JSON;
import lombok.Data;

import java.util.Map;

@Data
public class RetrieveChatData {
    private String id;
    private String conversation_id;
    private String bot_id;
    private String status;

    private Integer created_at;
    private Integer completed_at;
    private Integer failed_at;

    private Map<String,String> meta_data;
    private LastError last_error;

    private String section_id;
    private RequiredAction required_action;
    private Usage usage;


}
