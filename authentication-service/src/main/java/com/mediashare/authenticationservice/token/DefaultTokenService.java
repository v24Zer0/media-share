package com.mediashare.authenticationservice.token;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTCreationException;
import com.auth0.jwt.exceptions.JWTVerificationException;
import com.auth0.jwt.interfaces.Claim;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.mediashare.authenticationservice.util.UUIDRegex;
import org.springframework.stereotype.Service;

import java.util.Map;

@Service
public class DefaultTokenService implements TokenService {
    private final UUIDRegex uuidRegex;

    public DefaultTokenService() {
        this.uuidRegex = new UUIDRegex();
    }

    public String createToken(String userID) {
//        Check for valid userID (uuid or ksuid)
        if(!this.uuidRegex.validateID(userID)) {
            return "";
        }

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
            return false;
        }

//      Check user claim is present and non-empty
        Map<String, Claim> claims = decodedJWT.getClaims();
        if(!claims.containsKey("user"))
            return false;

        return !claims.get("user").asString().equals("");
    }
}
