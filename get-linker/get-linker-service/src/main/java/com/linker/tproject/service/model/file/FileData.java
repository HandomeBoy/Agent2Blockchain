package com.linker.tproject.service.model.file;

import lombok.Data;

@Data
public class FileData {
    private Integer bytes;
    private Integer created_at;
    private String file_name;
    private String id;
}
