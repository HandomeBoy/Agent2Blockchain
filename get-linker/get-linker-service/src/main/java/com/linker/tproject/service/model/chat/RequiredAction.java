package com.linker.tproject.service.model.chat;

import lombok.Data;

@Data
public class RequiredAction {
    private String type;
    private SubmitToolOutputs submitToolOutputs;
}
