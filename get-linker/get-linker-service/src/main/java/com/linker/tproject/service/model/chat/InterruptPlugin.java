package com.linker.tproject.service.model.chat;

import lombok.Data;

@Data
public class InterruptPlugin {
    private String id;
    private String type;
    private InterruptFunction interruptFunction;
}
