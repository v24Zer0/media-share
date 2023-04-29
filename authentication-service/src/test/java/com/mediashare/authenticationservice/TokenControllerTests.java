package com.mediashare.authenticationservice;

import com.mediashare.authenticationservice.exceptions.BadRequestException;
import com.mediashare.authenticationservice.exceptions.ForbiddenException;
import com.mediashare.authenticationservice.mock.MockTokenService;
import com.mediashare.authenticationservice.token.TokenController;
import com.mediashare.authenticationservice.token.TokenRequest;
import com.mediashare.authenticationservice.token.TokenResponse;
import com.mediashare.authenticationservice.token.TokenVerifyRequest;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.assertj.core.api.Assertions.assertThat;


public class TokenControllerTests {
    private TokenController tokenController;

    @BeforeEach
    public void setup() {
        this.tokenController = new TokenController(new MockTokenService());
    }

    @Test()
    @DisplayName("Should successfully create token and return valid TokenResponse")
    public void testCreateToken() {
        TokenResponse res = this.tokenController.createToken(new TokenRequest("user_01"));
        assertThat(res.getToken()).isNotEqualTo("");
    }

    @Test
    @DisplayName("Should fail to create token with invalid userID and throw BadRequestException")
    public void testCreateTokenWithBadRequest() {
        Exception e = assertThrows(BadRequestException.class, () ->
                this.tokenController.createToken(new TokenRequest("")));

        assertThat(e.getMessage()).isEqualTo("invalid userID");
    }

    @Test
    @DisplayName("Should fail to create token due to error from service")
    public void testCreateTokenFailedTokenCreation() {
        Exception e = assertThrows(BadRequestException.class, () ->
                this.tokenController.createToken(new TokenRequest("invalid_user")));

        assertThat(e.getMessage()).isEqualTo("could not create token");
    }

    @Test
    @DisplayName("Should successfully verify token")
    public void testVerifyToken() {
        this.tokenController.verifyToken(new TokenVerifyRequest("valid_token"));
    }

    @Test
    @DisplayName("Should fail to verify token with BadRequestException")
    public void testVerifyTokenWithBadRequestException() {
        Exception e = assertThrows(BadRequestException.class, () ->
                this.tokenController.verifyToken(new TokenVerifyRequest("")));

        assertThat(e.getMessage()).isEqualTo("invalid token");
    }

    @Test
    @DisplayName("Should fail to verify token with BadRequestException")
    public void testVerifyTokenWithForbiddenException() {
        Exception e = assertThrows(ForbiddenException.class, () ->
                this.tokenController.verifyToken(new TokenVerifyRequest("invalid_token")));

        assertThat(e.getMessage()).isEqualTo("could not verify token");
    }
}
