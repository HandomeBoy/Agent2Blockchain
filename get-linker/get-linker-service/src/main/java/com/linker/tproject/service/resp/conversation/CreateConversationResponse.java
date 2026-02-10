package com.linker.tproject.service.resp.conversation;

import com.linker.tproject.service.model.conversation.ConversationData;
import com.linker.tproject.service.model.conversation.ResponseDetail;
import lombok.Data;

@Data
public class CreateConversationResponse {
    private ConversationData data;
    private ResponseDetail detail;
    private Long code;
    private String msg;
}