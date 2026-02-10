package com.linker.tproject.service.model.conversation;

import lombok.Data;

import java.util.Map;

@Data
public class ConversationData {
    private String id;              // 会话唯一标识
    private String name;            // 会话名称
    private Map<String, Object> meta_data; // 附加信息，JSON Map 格式
    private String creator_id;      // 创建者扣子 UID
    private Long created_at;        // 创建时间（Unix 时间戳，秒）
    private Long updated_at;        // 最后更新时间（Unix 时间戳，秒）
    private String last_section_id; // 最新上下文片段 ID
    private String connector_id;    // 渠道 ID（如 API:1024, ChatSDK:999）
}
