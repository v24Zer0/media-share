package com.mediashare.authenticationservice;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.assertj.core.api.Assertions.assertThat;

public class TokenServiceTests {

    @Test
    @DisplayName("Should match given string")
    public void testMatchString() {
        String hello = "hello";
        assertThat(hello).isEqualTo("hello");
    }
}
