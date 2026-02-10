package com.linker.tproject.service.resp;

import lombok.Data;
import org.springframework.stereotype.Component;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

@Data
@Component
public class SetReturn {
    public String code;
    public String message;
    public Object data;
    public String success;
    public String timestamp;

    LocalDateTime now = LocalDateTime.now();
    String formattedTime = now.format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));

    public SetReturn() {
    }

    public SetReturn isSuccess(String message, Object data) {
        SetReturn result = new SetReturn();

        result.code = "0";
        result.message = message;
        result.data = data;
        result.success = "true";
        result.timestamp = formattedTime;
        return result;
    }

    public SetReturn isfail(String code, String message, Object data) {
        SetReturn result = new SetReturn();

        result.code = code;
        result.message = message;
        result.data = data;
        result.success = "false";
        result.timestamp = formattedTime;
        return result;
    }
}
