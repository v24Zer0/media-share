package com.mediashare.authenticationservice.token;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/token")
public class TokenController {
    private final TokenService tokenService;

    @Autowired
    public TokenController(TokenService tokenService) {
        this.tokenService = tokenService;
    }

    @PostMapping("/")
    public TokenResponse createToken(@RequestBody TokenRequest tokenRequest) {
//        verify RequestBody (Check that userID valid)
        if(tokenRequest.getUserID().equals("")) {
            return null;
        }

        String token = this.tokenService.createToken(tokenRequest.getUserID());
        return new TokenResponse(token);
    }

    @PostMapping("/verify")
    public void verifyToken(@RequestBody TokenVerifyRequest tokenVerifyRequest) {
        boolean res = this.tokenService.verifyToken(tokenVerifyRequest.getToken());
    }
}
