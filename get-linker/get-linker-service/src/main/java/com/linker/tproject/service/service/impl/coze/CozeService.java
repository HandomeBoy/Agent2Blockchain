package com.linker.tproject.service.service.impl.coze;

import com.linker.tproject.service.config.CozeConfig;
import com.linker.tproject.service.model.chat.GetAgentMessageData;
import com.linker.tproject.service.model.chat.Parameters;
import com.linker.tproject.service.req.chat.CreateChatRequest;
import com.linker.tproject.service.req.conversation.CreateConversationRequest;
import com.linker.tproject.service.req.message.CreateMessageRequest;
import com.linker.tproject.service.resp.chat.CreateChatResponse;
import com.linker.tproject.service.resp.chat.GetAgentMessageResponse;
import com.linker.tproject.service.resp.chat.GetRetrieveChatResponse;
import com.linker.tproject.service.resp.conversation.CreateConversationResponse;
import com.linker.tproject.service.resp.message.CreateMessageResponse;
import com.linker.tproject.service.resp.file.RetrieveFileResponse;
import com.linker.tproject.service.resp.file.UploadFileResponse;
import com.linker.tproject.service.service.impl.database.StoreAgentResults;
import com.linker.tproject.service.service.impl.file.TomatoImgService;
import lombok.extern.slf4j.Slf4j;
import lombok.var;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.FileSystemResource;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import java.io.File;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.Map;

import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.http.MediaType;

@Service
@Slf4j
public class CozeService {

    @Autowired
    private CozeConfig cozeConfig;

    @Autowired
    private RestTemplate restTemplate;

    @Autowired
    private StoreAgentResults storeAgentResults;

    @Autowired
    private TomatoImgService tomatoImgService;


    /**
     * 创建会话
     *
     * @return CreateConversationResponse 会话创建响应
     */
    public CreateConversationResponse createConversation() {
        CreateConversationRequest request = new CreateConversationRequest(
                cozeConfig.getCozeAgentId(), "番茄检测"
        );

        String url = cozeConfig.getCozeConservation();

        // 创建请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.set("Content-Type", "application/json");

        HttpEntity<CreateConversationRequest> entity = new HttpEntity<>(request, headers);

        try {
            // 发送请求
            ResponseEntity<CreateConversationResponse> response = restTemplate.exchange(
                    url, HttpMethod.POST, entity, CreateConversationResponse.class);
            log.info("创建会话成功,Coze API 响应: " + response.getBody());

            return response.getBody();
        } catch (RestClientException e) {
            log.error("调用 Coze API 失败: " + e.getMessage(), e);
            throw new RuntimeException(e);
        }
    }

    /**
     * 创建消息
     *
     * @param conversationId 会话 ID
     * @param fileId         文件 ID
     * @return CreateMessageResponse 消息创建响应
     */
    public CreateMessageResponse createMessage(String conversationId, String fileId) {

        // 构造请求 URL
        String url = cozeConfig.getCozeMessage() + "?conversation_id=" + conversationId;

        // 设置请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.set("Content-Type", "application/json");

        // 创建消息请求对象，配置默认值和传入的 fileId
        CreateMessageRequest messageRequest = new CreateMessageRequest();
        messageRequest.setRole("user");
        messageRequest.setContent_type("object_string");
        // 根据 API 要求构建内容数组，包含文本和图片
        String content = "[{\"type\":\"text\",\"text\":\"番茄检测。\"},{\"type\":\"image\",\"file_id\":\"" + fileId + "\"}]";
        messageRequest.setContent(content);

        // 创建请求实体
        HttpEntity<CreateMessageRequest> entity = new HttpEntity<>(messageRequest, headers);

        // 发送 POST 请求
        ResponseEntity<CreateMessageResponse> response = restTemplate.exchange(
                url,
                HttpMethod.POST,
                entity,
                CreateMessageResponse.class
        );

        return response.getBody();
    }


    /**
     * 上传文件
     *
     * @param filename 文件名
     * @return UploadFileResponse 文件上传响应
     */
    public UploadFileResponse uploadFile(String filename) {
        // 构造文件路径
        Path filePath = Paths.get(cozeConfig.getImgRootPath(), filename + ".jpg").normalize();
        File file = filePath.toFile();

        log.info("完整文件路径: {}", file.getAbsolutePath());

        // 验证文件是否存在
        if (!file.exists()) {
            throw new IllegalArgumentException("文件不存在: " + file.getAbsolutePath());
        }

        // 创建 MultiValueMap 表单数据
        MultiValueMap<String, Object> formData = new LinkedMultiValueMap<>();
        formData.add("file", new FileSystemResource(file)); // 使用 FileSystemResource 上传文件

        // 设置请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.setContentType(MediaType.MULTIPART_FORM_DATA); // 设置为 multipart/form-data

        // 创建请求实体
        HttpEntity<MultiValueMap<String, Object>> entity = new HttpEntity<>(formData, headers);

        // 发送 POST 请求
        ResponseEntity<UploadFileResponse> response = restTemplate.exchange(
                cozeConfig.getCozeUploadFile(),
                HttpMethod.POST,
                entity,
                UploadFileResponse.class
        );

        return response.getBody();
    }

    /**
     * 获取文件详情
     *
     * @param fileId 文件 ID
     * @return RetrieveFileResponse 文件详情响应
     */
    public RetrieveFileResponse retrieveFile(String fileId) {
        // 构造请求 URL，添加 file_id 查询参数
        String url = cozeConfig.getCozeGetFile() + "?file_id=" + fileId;

        // 创建请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.set("Content-Type", "application/json");

        // 创建请求实体
        HttpEntity<Void> entity = new HttpEntity<>(headers);

        try {
            // 发送 GET 请求，返回类型为 RetrieveFileResponse
            ResponseEntity<RetrieveFileResponse> response = restTemplate.exchange(
                    url,
                    HttpMethod.GET,
                    entity,
                    RetrieveFileResponse.class
            );

            return response.getBody();
        } catch (RestClientException e) {
            log.error("调用 Coze 获取文件详情失败: {}", e.getMessage(), e);
            throw new RuntimeException("获取文件详情失败", e);
        }
    }


    public CreateChatResponse createChat(String conversation_id, String fileId) {

        log.info("创建对话");
        log.info("文件ID: {}", fileId);

        String url = cozeConfig.getCozeChat() + "?conversation_id=" + conversation_id;
        CreateChatRequest chatRequest = new CreateChatRequest();
        Parameters parameters = new Parameters();
        Map<String, String> imageInfo = new HashMap<>();
        imageInfo.put("file_id", fileId);
        parameters.setImage(String.valueOf(imageInfo));
        // 创建请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.set("Content-Type", "application/json");

        chatRequest.setUser_id("123");
        chatRequest.setBot_id(cozeConfig.getCozeAgentId());
        chatRequest.setStream(false);
        chatRequest.setAuto_save_history(true);
        chatRequest.setParameters(parameters);


        // 创建请求实体
        HttpEntity<CreateChatRequest> entity = new HttpEntity<>(chatRequest, headers);

        log.info("请求: {}", entity);

        try {
            // 发送 Post 请求，返回类型为 CreateChatResponse
            ResponseEntity<CreateChatResponse> response = restTemplate.exchange(
                    url,
                    HttpMethod.POST,
                    entity,
                    CreateChatResponse.class
            );

            return response.getBody();
        } catch (RestClientException e) {
            log.error("调用 Coze 创建对话失败: {}", e.getMessage(), e);
            throw new RuntimeException("创建对话失败", e);
        }
    }

    public GetAgentMessageResponse getAgentMessage(String conversationId, String chatId) {
        log.info("获取对话消息");
        log.info("会话ID: {}", conversationId);
        log.info("对话ID: {}", chatId);

        // 构建请求URL，添加conversation_id和chat_id查询参数
        String url = cozeConfig.getCozeGetAgentMessageList() +
                "?conversation_id=" + conversationId +
                "&chat_id=" + chatId;

        // 创建请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.set("Content-Type", "application/json");

        // 创建请求实体
        HttpEntity<Void> entity = new HttpEntity<>(headers);

        try {
            // 发送 GET 请求，返回类型为 GetAgentMessageResponse
            ResponseEntity<GetAgentMessageResponse> response = restTemplate.exchange(
                    url,
                    HttpMethod.GET,
                    entity,
                    GetAgentMessageResponse.class
            );

            GetAgentMessageResponse responseBody = response.getBody();
            log.info("获取对话消息成功: {}", responseBody);

            // 安全访问响应数据
            if (responseBody != null && responseBody.getData() != null && responseBody.getData().size() > 1) {
                if (responseBody.getData().get(1) != null && responseBody.getData().get(1).getContent() != null) {
                    log.info("智能体分析结果: {}", responseBody.getData().get(1).getContent());
                } else {
                    log.warn("响应数据中第2个元素或其内容为空");
                }
            } else {
                log.warn("响应数据为空或数据列表长度不足");
            }

            return responseBody;
        } catch (RestClientException e) {
            log.error("调用 Coze 获取对话消息失败: {}", e.getMessage(), e);
            throw new RuntimeException("获取对话消息失败", e);
        }
    }

    public GetRetrieveChatResponse waitForAgentResponse(String conversationId, String chatId) {
        log.info("等待智能体响应");
        String url = cozeConfig.getCozeRetrieveChat() + "?conversation_id=" + conversationId + "&chat_id=" + chatId;
        log.info("请求: {}", url);
        // 创建请求头
        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer " + cozeConfig.getCozeBearer());
        headers.set("Content-Type", "application/json");

        // 创建请求实体
        HttpEntity<Void> entity = new HttpEntity<>(headers);
        try {
            ResponseEntity<GetRetrieveChatResponse> response = restTemplate.exchange(
                    url,
                    HttpMethod.GET,
                    entity,
                    GetRetrieveChatResponse.class
            );

            return response.getBody();
        } catch (RestClientException e) {
            throw new RuntimeException(e);
        }
    }


    private GetRetrieveChatResponse pollForAgentResponse(String conversationId, String chatId) {
        final int POLL_INTERVAL = 1000; // 每秒轮询一次（根据API文档建议）
        final int TIMEOUT = 120000; // 2分钟超时（毫秒）
        long startTime = System.currentTimeMillis();

        while (System.currentTimeMillis() - startTime < TIMEOUT) {
            try {
                GetRetrieveChatResponse response = waitForAgentResponse(conversationId, chatId);

                // 检查是否获取到有效结果
                if (isChatCompleted(response)) {
                    log.info("智能体分析完成");
                    return response;
                }

                log.info("等待智能体分析完成，已等待: {} 秒",
                        (System.currentTimeMillis() - startTime) / 1000);

                Thread.sleep(POLL_INTERVAL);
            } catch (InterruptedException e) {
                log.error("轮询被中断", e);
                Thread.currentThread().interrupt();
                break;
            } catch (Exception e) {
                log.error("轮询智能体响应时发生错误", e);
                // 发生错误时也继续轮询，直到超时
                try {
                    Thread.sleep(POLL_INTERVAL);
                } catch (InterruptedException ie) {
                    Thread.currentThread().interrupt();
                    break;
                }
            }
        }

        log.warn("等待智能体响应超时（3分钟）");
        return null;
    }

    /**
     * 检查对话是否已完成
     *
     * @param response 响应对象
     * @return 是否已完成
     */
    private boolean isChatCompleted(GetRetrieveChatResponse response) {
        if (response == null || response.getData() == null) {
            return false;
        }

        String status = response.getData().getStatus();
        return "completed".equals(status); // 确保状态值完全匹配
    }

    // 根据用户身份自动执行番茄检测,harvest:收获,sampling:采样
    // 其余用户角色无法执行番茄检测
    public boolean autoDetect(String tomato_id,String role) {

        if (!role.equals("harvest") && !role.equals("sampling")) {
            log.info("用户没有执行番茄检测的权限");
            return false;
        }

        //本地化存储图片
        String fileName = tomatoImgService.saveHarvestImageToFile(tomato_id,role);
        log.info("文件名: {}", fileName);
        // 移除文件后缀，因为 uploadFile 方法会自动添加 .jpg 后缀
        String fileNameWithoutExt;
        int lastDotIndex = fileName.lastIndexOf('.');
        if (lastDotIndex > 0) {
            fileNameWithoutExt = fileName.substring(0, lastDotIndex);
        } else {
            fileNameWithoutExt = fileName; // 如果没有后缀则直接使用原名
        }

        //上传文件,获取文件ID
        UploadFileResponse uploadFileResponse = uploadFile(fileNameWithoutExt);
        if (uploadFileResponse.getData() == null) {
            log.error("上传文件响应数据为空，文件名: {}", fileNameWithoutExt);
            return false;
        }

        String fileId = uploadFileResponse.getData().getId();
        if (fileId == null) {
            log.error("获取到的文件ID为空，文件名: {}", fileNameWithoutExt);
            return false;
        }

        log.info("文件ID: {}", fileId);
        //创建会话,获取会话ID
        CreateConversationResponse createConversationResponse = createConversation();
        String conversationId = createConversationResponse.getData().getId();
        //创建消息,获取消息ID
        CreateMessageResponse createMessageResponse = createMessage(conversationId, fileId);
        //创建对话,传入文件,获取智能体分析结果,并存入数据库
        //创建对话
        CreateChatResponse createChatResponse = createChat(conversationId, fileId);
        if (createChatResponse == null || createChatResponse.getData() == null) {
            throw new RuntimeException("创建对话失败");
        }

        String chatId = createChatResponse.getData().getId();
        log.info("创建对话结果: {}", createChatResponse.getData());

        // 等待智能体分析结果，轮询直到完成或超时
        GetRetrieveChatResponse retrieveChatResponse = pollForAgentResponse(conversationId, chatId);

        //获取智能体分析结果
        GetAgentMessageResponse getAgentMessageResponse = getAgentMessage(conversationId, chatId);
        if (getAgentMessageResponse == null || getAgentMessageResponse.getData() == null) {
            throw new RuntimeException("获取智能体分析结果失败，响应数据为空");
        }

        // 提取分析结果
        String agentMessage = getAgentMessageResponse.getData().get(1).getContent();

        // 执行存储操作
        if (role.equals("harvest")) {
            return storeAgentResults.storeHarvestResult(tomato_id, agentMessage);
        } else {
            return storeAgentResults.storeSamplingResult(tomato_id, agentMessage);
        }
    }
}