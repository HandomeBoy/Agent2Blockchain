package com.linker.tproject.service.req.chat;

import com.linker.tproject.service.model.chat.Parameters;
import lombok.Data;

/*
* "user_id": "123",
  "stream": false,
  "auto_save_history": true,
  "parameters": {
    "img": "{\"file_id\":\"file_id\"}"
  }
*/
@Data
public class CreateChatRequest {
    private String user_id;
    private String bot_id;
    private boolean stream;
    private boolean auto_save_history;
    private Parameters parameters;



}
