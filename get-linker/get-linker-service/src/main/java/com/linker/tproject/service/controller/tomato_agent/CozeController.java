package com.linker.tproject.service.controller.tomato_agent;

import com.linker.tproject.service.resp.SetReturn;
import com.linker.tproject.service.resp.chat.CreateChatResponse;
import com.linker.tproject.service.resp.chat.GetAgentMessageResponse;
import com.linker.tproject.service.resp.chat.GetRetrieveChatResponse;
import com.linker.tproject.service.resp.conversation.CreateConversationResponse;
import com.linker.tproject.service.resp.message.CreateMessageResponse;
import com.linker.tproject.service.resp.file.RetrieveFileResponse;
import com.linker.tproject.service.resp.file.UploadFileResponse;
import com.linker.tproject.service.service.impl.coze.CozeService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@Api(tags = "Coze服务")
public class CozeController {

    @Autowired
    private CozeService cozeService;
    @Autowired
    private SetReturn setReturn;

    /**
     * 自动进行检测并存入数据库
     * @param tomatoId
     * @return
     */
    @PostMapping("/tomato-detect")
    @ApiOperation(value = "番茄检测")
    public SetReturn tomatoDetect(@RequestParam String tomatoId,@RequestParam String role) {
        if (cozeService.autoDetect(tomatoId, role)){
            return setReturn.isSuccess("成功", "");
        }else {
            return setReturn.isfail("200", "失败", "");
        }
    }

    @PostMapping("/create-conversation")
    @ApiOperation(value = "创建会话")
    public CreateConversationResponse createConversation() {
        return cozeService.createConversation();
    }

    @PostMapping("/create-message")
    @ApiOperation(value = "创建消息")
    public CreateMessageResponse createMessage(@RequestParam String conversationId, @RequestParam String fileId) {
        return cozeService.createMessage(conversationId, fileId);
    }

    @PostMapping("/upload-file")
    @ApiOperation(value = "上传文件")
    public UploadFileResponse uploadFile(@RequestParam String tomatoId) {
        return cozeService.uploadFile(tomatoId);
    }

    @PostMapping("/retrieve-file")
    @ApiOperation(value = "获取文件")
    public RetrieveFileResponse retrieveFile(@RequestParam String fileId) {
        return cozeService.retrieveFile(fileId);
    }

    @PostMapping("/create-chat")
    @ApiOperation(value = "创建对话")
    public CreateChatResponse CreateChat(@RequestParam String conversationId,@RequestParam String fileId) {

        return cozeService.createChat(conversationId,fileId);
    }

    @GetMapping("/get-agent-message")
    @ApiOperation(value = "获取智能体分析结果")
    public GetAgentMessageResponse getAgentMessage(@RequestParam String conversationId,@RequestParam String chatId) {
        return cozeService.getAgentMessage(conversationId,chatId);
    }

    @GetMapping("/wait-for-agent-response")
    @ApiOperation(value = "等待智能体回复")
    public GetRetrieveChatResponse waitForAgentResponse(@RequestParam String conversationId,@RequestParam String chatId) {
        return cozeService.waitForAgentResponse(conversationId,chatId);
    }

}
