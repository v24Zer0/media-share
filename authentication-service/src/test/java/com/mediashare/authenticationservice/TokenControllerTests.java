package com.mediashare.authenticationservice;

import com.mediashare.authenticationservice.token.TokenController;
import com.mediashare.authenticationservice.token.TokenRequest;
import com.mediashare.authenticationservice.token.TokenResponse;
import com.mediashare.authenticationservice.token.TokenService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class TokenControllerTests {
    TokenController tokenController;

    @BeforeEach
    public void setup() {
        this.tokenController = new TokenController(null);
    }

    @Test
    @DisplayName("Should create token")
    public void testCreateToken() {
//        TokenResponse res = this.tokenController.createToken(new TokenRequest(""));
    }
}
