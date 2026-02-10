package com.linker.tproject.service.resp.chat;

import com.linker.tproject.service.model.chat.RetrieveChatData;
import com.linker.tproject.service.model.chat.RetrieveChatDetail;
import lombok.Data;

@Data
public class GetRetrieveChatResponse {
    private RetrieveChatData data;
    private RetrieveChatDetail detail;
    private Integer code;
    private String message;
}
