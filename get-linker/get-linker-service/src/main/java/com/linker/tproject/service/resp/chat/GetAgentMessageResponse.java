package com.linker.tproject.service.resp.chat;

import com.linker.tproject.service.model.chat.GetAgentMessageData;
import com.linker.tproject.service.model.chat.GetAgentMessageDetail;
import io.swagger.models.auth.In;
import lombok.Data;

import java.util.List;

@Data
public class GetAgentMessageResponse {
    private Integer code;
    private List<GetAgentMessageData> data;
    private GetAgentMessageDetail detail;
    private String msg;
}
