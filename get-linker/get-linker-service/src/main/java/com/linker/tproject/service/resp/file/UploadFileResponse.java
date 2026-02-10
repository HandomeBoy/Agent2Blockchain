package com.linker.tproject.service.resp.file;



import com.linker.tproject.service.model.file.FileData;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class UploadFileResponse {
    private int code;
    private String msg;
    private FileData data;
}
