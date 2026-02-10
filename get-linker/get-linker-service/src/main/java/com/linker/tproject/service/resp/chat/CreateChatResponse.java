package com.linker.tproject.service.resp.chat;

import com.linker.tproject.service.model.chat.CreateChatData;
import com.linker.tproject.service.model.chat.CreateChatDetail;
import lombok.Data;

@Data
public class CreateChatResponse {
    private Integer code;
    private CreateChatData data;
    private CreateChatDetail detail;
    private String msg;
}
