package com.linker.tproject.service.resp.message;

import com.linker.tproject.service.model.message.MessageData;
import com.linker.tproject.service.model.message.MessageDetail;
import lombok.Data;

@Data
public class CreateMessageResponse {
    private Integer code;
    private MessageData data;
    private MessageDetail detail;
    private String msg;
}
