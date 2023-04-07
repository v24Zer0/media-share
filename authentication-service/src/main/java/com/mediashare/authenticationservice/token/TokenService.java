package com.mediashare.authenticationservice.token;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTCreationException;
import org.springframework.stereotype.Service;

@Service
public class TokenService {

    public String getToken() {
        try {
            Algorithm algorithm = Algorithm.HMAC256("secret");
            return JWT.create().withClaim("user", "userID").sign(algorithm);
        } catch (JWTCreationException exception) {
            return "";
        }
    }
}
