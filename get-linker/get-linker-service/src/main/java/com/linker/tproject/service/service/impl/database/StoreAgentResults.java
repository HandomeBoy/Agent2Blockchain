package com.linker.tproject.service.service.impl.database;

import com.linker.tproject.service.entity.TomatoInfo;
import com.linker.tproject.service.mapper.TomatoInfoMapper;
import lombok.Data;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

@Data
@Service
public class StoreAgentResults {

    @Autowired
    TomatoInfoMapper tomatoInfoMapper;

    //Mybatis-plus存储result数据
    public boolean storeHarvestResult(String tomatoId, String result) {
        // 创建实体对象
        TomatoInfo tomatoInfo = new TomatoInfo();
        tomatoInfo.setTomatoId(tomatoId);  // 设置ID
        tomatoInfo.setHarvestEvaluation(result); // 设置要更新的字段

        // 执行更新操作
        int rows = tomatoInfoMapper.updateById(tomatoInfo);

        // 返回更新结果
        return rows > 0;
    }

    public boolean storeSamplingResult(String tomatoId, String result) {
        // 创建实体对象
        TomatoInfo tomatoInfo = new TomatoInfo();
        tomatoInfo.setTomatoId(tomatoId);  // 设置ID
        tomatoInfo.setSamplingEvaluation(result); // 设置要更新的字段

        // 执行更新操作
        int rows = tomatoInfoMapper.updateById(tomatoInfo);

        // 返回更新结果
        return rows > 0;
    }
}
