package com.linker.tproject.service.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.springframework.context.annotation.Bean;

@Data
@TableName("tomato_info")
public class TomatoInfo {
    @TableId(value = "tomato_id")
    private String tomatoId;

    private String productionArea;
    private java.sql.Date harvestDate;
    private java.sql.Date processingDate;
    private String transportInfo;
    private byte[] harvestImage;
    private byte[] samplingImage;
    private java.sql.Timestamp createdAt;
    private java.sql.Timestamp updatedAt;
    private String harvestEvaluation;
    private String samplingEvaluation;
}