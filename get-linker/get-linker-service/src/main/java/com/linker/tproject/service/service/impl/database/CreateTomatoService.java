package com.linker.tproject.service.service.impl.database;

import com.linker.tproject.service.entity.TomatoInfo;
import com.linker.tproject.service.mapper.TomatoInfoMapper;
import com.linker.tproject.service.service.ICreateTomato;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Repository;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.sql.Date;
import java.sql.Timestamp;

@Repository
@Slf4j
public class CreateTomatoService implements ICreateTomato {

    @Autowired
    TomatoInfoMapper tomatoInfoMapper;

    @Value("${tomato.harvest.image.path}")
    private String harvestBasePath;

    @Value("${tomato.sampling.image.path}")
    private String samplingBasePath;

    @Override
    public Boolean createTomato(String tomato_id,
                                String production_area,
                                Date harvest_date,
                                Date processing_date,
                                String transport_info,
                                Timestamp created_at) {

        // 构建完整文件路径
        String harvestImagePath = Paths.get(harvestBasePath, tomato_id + ".jpg").toString();
        String samplingImagePath = Paths.get(samplingBasePath, tomato_id + ".jpg").toString();

        byte[] harvest_image;
        byte[] sampling_image;

        try {
            harvest_image = Files.readAllBytes(Paths.get(harvestImagePath));
        } catch (IOException e) {
            throw new RuntimeException("读取收获图像失败: " + e.getMessage(), e);
        }

        try {
            sampling_image = Files.readAllBytes(Paths.get(samplingImagePath));
        } catch (IOException e) {
            throw new RuntimeException("读取采样图像失败: " + e.getMessage(), e);
        }

        // 创建 TomatoInfo 对象并设置属性
        TomatoInfo tomatoInfo = new TomatoInfo();
        tomatoInfo.setTomatoId(tomato_id);
        tomatoInfo.setCreatedAt(created_at);
        tomatoInfo.setTransportInfo(transport_info);
        tomatoInfo.setUpdatedAt(created_at);
        tomatoInfo.setHarvestDate(harvest_date);
        tomatoInfo.setProcessingDate(processing_date);
        tomatoInfo.setHarvestImage(harvest_image);
        tomatoInfo.setSamplingImage(sampling_image);
        tomatoInfo.setProductionArea(production_area);
        tomatoInfoMapper.insert(tomatoInfo);

        return true;

    }
}
