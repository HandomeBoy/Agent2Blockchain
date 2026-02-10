package com.linker.tproject.service.service.impl.file;

import com.linker.tproject.service.entity.TomatoInfo;
import com.linker.tproject.service.mapper.TomatoInfoMapper;
import com.linker.tproject.service.service.IGetTomatoImg;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

@Service
@Slf4j
public class TomatoImgService implements IGetTomatoImg {

    @Value("${tomato.agent.img.path}")
    private String outputPath;

    @Autowired
    private TomatoInfoMapper tomatoInfoMapper;

    @Override
    public String saveHarvestImageToFile(String tomatoId,String role) {
        log.info("保存收获图片到文件");
        log.info("OutputPath: " + outputPath);
        log.info("TomatoId: " + tomatoId);

        Path absolutePath = Paths.get(outputPath).toAbsolutePath();
        log.info("绝对路径: " + absolutePath.toString());

        // 验证输出路径是否有效
        if (outputPath == null || outputPath.trim().isEmpty()) {
            log.error("输出路径未配置或为空");
            return null;
        }

        TomatoInfo tomato = tomatoInfoMapper.selectById(tomatoId);

        if (tomato == null) {
            log.warn("番茄信息不存在: " + tomatoId);
            return null;
        }

        byte[] harvestImage = tomato.getHarvestImage();
        if (harvestImage == null || harvestImage.length == 0) {
            log.warn("收获图片不存在或为空: " + tomatoId);
            return null;
        }

        try {
            String fileName = tomatoId + "_" + role + ".jpg";
            Path outputDir = Paths.get(outputPath);

            // 确保目录存在且可写
            if (!Files.exists(outputDir)) {
                Files.createDirectories(outputDir);
                log.info("创建输出目录: " + outputDir.toString());
            }

            if (!Files.isWritable(outputDir)) {
                log.error("输出目录不可写: " + outputDir.toString());
                return null;
            }

            Files.write(Paths.get(outputPath, fileName), harvestImage);
            log.info("收获图片保存成功，文件名: " + fileName);
            return fileName;
        } catch (IOException e) {
            log.error("保存收获图像失败: " + e.getMessage(), e);
            return null;
        }
    }

}
