package com.linker.tproject.service.model.chat;

import lombok.Data;

@Data
public class Usage {
    private Integer input_count;
    private Integer token_count;
    private Integer output_count;
}
