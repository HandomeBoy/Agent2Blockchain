package com.linker.tproject.service.resp.file;

import com.linker.tproject.service.model.conversation.ConversationData;
import com.linker.tproject.service.model.conversation.ResponseDetail;
import com.linker.tproject.service.model.file.FileData;
import lombok.Data;

@Data
public class RetrieveFileResponse {
    private FileData data;
    private ResponseDetail detail;
    private Long code;
    private String msg;
}
