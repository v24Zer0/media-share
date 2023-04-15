package com.mediashare.authenticationservice.token;

import com.mediashare.authenticationservice.exceptions.BadRequestException;
import com.mediashare.authenticationservice.exceptions.ForbiddenException;
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

    @PostMapping("")
    public TokenResponse createToken(@RequestBody TokenRequest tokenRequest) {
        if(tokenRequest.getUserID().equals("")) {
            throw new BadRequestException("invalid userID");
        }

        String token = this.tokenService.createToken(tokenRequest.getUserID());
        if(token.equals("")) {
            throw new BadRequestException("could not create token");
        }

        return new TokenResponse(token);
    }

    @PostMapping("/verify")
    public void verifyToken(@RequestBody TokenVerifyRequest tokenVerifyRequest) {
        if(tokenVerifyRequest.getToken().equals("")) {
            throw new BadRequestException("invalid token");
        }

        boolean ok = this.tokenService.verifyToken(tokenVerifyRequest.getToken());
        if(!ok) {
            throw new ForbiddenException("could not verify token");
        }
    }
}
