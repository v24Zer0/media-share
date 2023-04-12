package com.mediashare.authenticationservice.token;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTCreationException;
import com.auth0.jwt.exceptions.JWTVerificationException;
import com.auth0.jwt.interfaces.DecodedJWT;
import org.springframework.stereotype.Service;

@Service
public class DefaultTokenService implements TokenService {

    public String createToken(String userID) {
        try {
            Algorithm algorithm = Algorithm.HMAC256("secret");
            return JWT.create().withClaim("user", userID).sign(algorithm);
        } catch (JWTCreationException exception) {
            return "";
        }
    }

    public boolean verifyToken(String token) {
        Algorithm algorithm = Algorithm.HMAC256("secret");
        JWTVerifier verifier = JWT.require(algorithm).build();
        DecodedJWT decodedJWT;

        try {
            decodedJWT = verifier.verify(token);
        } catch(JWTVerificationException exception) {
            System.out.println("Failed to verify token");
            return false;
        }

//      Check user claim is present and non-empty
        if(decodedJWT.getClaim("user").asString().equals("")) {
            return false;
        }
        return true;
    }
}
