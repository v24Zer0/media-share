package com.mediashare.authenticationservice;

import com.mediashare.authenticationservice.token.DefaultTokenService;
import com.mediashare.authenticationservice.token.TokenService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;

public class TokenServiceTests {
    private TokenService tokenService;
    @BeforeEach
    public void setup() {
        this.tokenService = new DefaultTokenService();
    }

    @Test
    @DisplayName("Should successfully create token with user claim matching given userID")
    public void testCreateToken() {
        String token = this.tokenService.createToken("user_01");
        assertThat(token).isNotEqualTo("");
    }

    @Test
    @DisplayName("Should fail to create token with JWTException thrown")
    public void testCreateTokenWithJWTCreationException() {

    }

    @Test
    @DisplayName("Should successfully verify token")
    public void testVerifyToken() {

    }

    @Test
    @DisplayName("Should fail to verify token with JWTVerificationException")
    public void testVerifyTokenWithJWTVerificationException() {

    }

    @Test
    @DisplayName("Should fail to verify token because of missing user claim")
    public void testVerifyTokenWithMissingUserClaim() {

    }
}
